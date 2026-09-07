
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RunMetricService } from '../../../services/RunMetric.service';
import { RunMetric } from '../../../models/RunMetric';

@Component({
    selector: 'app-index-runMetric',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRunMetricComponent implements OnInit {

    runMetrics: RunMetric[] = [];

    constructor(
        private router: Router,
        private service: RunMetricService
) {}

    ngOnInit(): void {
        this.getRunMetrics();
}

    getRunMetrics(): void {
        this.service.getRunMetrics().subscribe((res) => {
        this.runMetrics = res;
    });
}

    deleteRunMetric(id: any): void {
        this.service.deleteRunMetric(id)
            .subscribe(() => {
                this.getRunMetrics();
            });
    }
}