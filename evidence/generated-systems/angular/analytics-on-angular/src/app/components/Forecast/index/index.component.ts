
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ForecastService } from '../../../services/Forecast.service';
import { Forecast } from '../../../models/Forecast';

@Component({
    selector: 'app-index-forecast',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexForecastComponent implements OnInit {

    forecasts: Forecast[] = [];

    constructor(
        private router: Router,
        private service: ForecastService
) {}

    ngOnInit(): void {
        this.getForecasts();
}

    getForecasts(): void {
        this.service.getForecasts().subscribe((res) => {
        this.forecasts = res;
    });
}

    deleteForecast(id: any): void {
        this.service.deleteForecast(id)
            .subscribe(() => {
                this.getForecasts();
            });
    }
}