
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PerformanceMetricService } from '../../../services/PerformanceMetric.service';
import { PerformanceMetric } from '../../../models/PerformanceMetric';

@Component({
    selector: 'app-index-performanceMetric',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPerformanceMetricComponent implements OnInit {

    performanceMetrics: PerformanceMetric[] = [];

    constructor(
        private router: Router,
        private service: PerformanceMetricService
) {}

    ngOnInit(): void {
        this.getPerformanceMetrics();
}

    getPerformanceMetrics(): void {
        this.service.getPerformanceMetrics().subscribe((res) => {
        this.performanceMetrics = res;
    });
}

    deletePerformanceMetric(id: any): void {
        this.service.deletePerformanceMetric(id)
            .subscribe(() => {
                this.getPerformanceMetrics();
            });
    }
}