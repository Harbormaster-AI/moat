import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ForecastService } from '../../../services/Forecast.service';
import { SubBaseComponent } from '../../Forecast/sub.base.component';


@Component({
    selector: 'app-edit-forecast',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditForecastComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Forecast';

    forecastForm: FormGroup;
    forecast: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ForecastService,
        private fb: FormBuilder
) {
        super(http);
        this.forecastForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  forecastNumber: ['', Validators.required],
      forecastHorizonStart: ['', Validators.required],
      forecastHorizonEnd: ['', Validators.required],
      Lines: ['', ],
      Method: ['', ]
        });
    }

    
    updateForecast(forecastNumber, forecastHorizonStart, forecastHorizonEnd, Lines, Method): void {
        this.route.params.subscribe((params) => {

                        this.service.updateForecast(forecastNumber, forecastHorizonStart, forecastHorizonEnd, Lines, Method, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexForecast']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getForecast(params['id']).subscribe(res => {
                this.forecast = res;
            });
        });
    }
}