import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ForecastLineService } from '../../../services/ForecastLine.service';
import { SubBaseComponent } from '../../ForecastLine/sub.base.component';


@Component({
    selector: 'app-edit-forecastLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditForecastLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ForecastLine';

    forecastLineForm: FormGroup;
    forecastLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ForecastLineService,
        private fb: FormBuilder
) {
        super(http);
        this.forecastLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  period: ['', Validators.required],
      quantity: ['', Validators.required],
      confidence: ['', Validators.required],
      Forecast: ['', ],
      Item: ['', ]
        });
    }

    
    updateForecastLine(period, quantity, confidence, Forecast, Item): void {
        this.route.params.subscribe((params) => {

                        this.service.updateForecastLine(period, quantity, confidence, Forecast, Item, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexForecastLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getForecastLine(params['id']).subscribe(res => {
                this.forecastLine = res;
            });
        });
    }
}