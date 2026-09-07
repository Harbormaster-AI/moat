
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ForecastLineService } from '../../../services/ForecastLine.service';
import { ForecastLine } from '../../../models/ForecastLine';

@Component({
    selector: 'app-index-forecastLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexForecastLineComponent implements OnInit {

    forecastLines: ForecastLine[] = [];

    constructor(
        private router: Router,
        private service: ForecastLineService
) {}

    ngOnInit(): void {
        this.getForecastLines();
}

    getForecastLines(): void {
        this.service.getForecastLines().subscribe((res) => {
        this.forecastLines = res;
    });
}

    deleteForecastLine(id: any): void {
        this.service.deleteForecastLine(id)
            .subscribe(() => {
                this.getForecastLines();
            });
    }
}