import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ForecastService } from '../../../services/Forecast.service';
import { Forecast } from '../../../models/Forecast';
import { SubBaseComponent } from '../../Forecast/sub.base.component';

@Component({
    selector: 'app-create-forecast',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateForecastComponent extends SubBaseComponent implements OnInit {

    title = 'Add Forecast';

    forecastForm: FormGroup;
    forecast: Forecast;

    constructor( http: HttpClient,
        private forecastService: ForecastService,
        private fb: FormBuilder,
        private router: Router
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

    
    addForecast(forecastNumber, forecastHorizonStart, forecastHorizonEnd, Lines, Method): void {
        this.forecastService
        .addForecast(forecastNumber, forecastHorizonStart, forecastHorizonEnd, Lines, Method)
            .subscribe(() => {
                this.router.navigate(['/indexForecast']);
            });
    }

    ngOnInit(): void {
    }
}