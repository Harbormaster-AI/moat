
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MetricService } from '../../../services/Metric.service';
import { Metric } from '../../../models/Metric';

@Component({
    selector: 'app-index-metric',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMetricComponent implements OnInit {

    metrics: Metric[] = [];

    constructor(
        private router: Router,
        private service: MetricService
) {}

    ngOnInit(): void {
        this.getMetrics();
}

    getMetrics(): void {
        this.service.getMetrics().subscribe((res) => {
        this.metrics = res;
    });
}

    deleteMetric(id: any): void {
        this.service.deleteMetric(id)
            .subscribe(() => {
                this.getMetrics();
            });
    }
}