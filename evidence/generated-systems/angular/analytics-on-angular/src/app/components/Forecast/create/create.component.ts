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
                  name: ['', Validators.required],
      horizon: ['', Validators.required],
      ModelVersion: ['', ],
      TimeSeries: ['', ],
      Datasets: ['', ],
      Granularity: ['', ]
        });
    }

    
    addForecast(name, horizon, ModelVersion, TimeSeries, Datasets, Granularity): void {
        this.forecastService
        .addForecast(name, horizon, ModelVersion, TimeSeries, Datasets, Granularity)
            .subscribe(() => {
                this.router.navigate(['/indexForecast']);
            });
    }

    ngOnInit(): void {
    }
}