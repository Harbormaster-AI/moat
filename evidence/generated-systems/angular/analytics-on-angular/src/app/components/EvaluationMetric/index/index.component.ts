
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EvaluationMetricService } from '../../../services/EvaluationMetric.service';
import { EvaluationMetric } from '../../../models/EvaluationMetric';

@Component({
    selector: 'app-index-evaluationMetric',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEvaluationMetricComponent implements OnInit {

    evaluationMetrics: EvaluationMetric[] = [];

    constructor(
        private router: Router,
        private service: EvaluationMetricService
) {}

    ngOnInit(): void {
        this.getEvaluationMetrics();
}

    getEvaluationMetrics(): void {
        this.service.getEvaluationMetrics().subscribe((res) => {
        this.evaluationMetrics = res;
    });
}

    deleteEvaluationMetric(id: any): void {
        this.service.deleteEvaluationMetric(id)
            .subscribe(() => {
                this.getEvaluationMetrics();
            });
    }
}