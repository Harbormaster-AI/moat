import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EvaluationMetricService } from '../../../services/EvaluationMetric.service';
import { EvaluationMetric } from '../../../models/EvaluationMetric';
import { SubBaseComponent } from '../../EvaluationMetric/sub.base.component';

@Component({
    selector: 'app-create-evaluationMetric',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEvaluationMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Add EvaluationMetric';

    evaluationMetricForm: FormGroup;
    evaluationMetric: EvaluationMetric;

    constructor( http: HttpClient,
        private evaluationMetricService: EvaluationMetricService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.evaluationMetricForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      value: ['', Validators.required],
      ModelVersion: ['', ],
      Metric: ['', ],
      Dataset: ['', ]
        });
    }

    
    addEvaluationMetric(name, value, ModelVersion, Metric, Dataset): void {
        this.evaluationMetricService
        .addEvaluationMetric(name, value, ModelVersion, Metric, Dataset)
            .subscribe(() => {
                this.router.navigate(['/indexEvaluationMetric']);
            });
    }

    ngOnInit(): void {
    }
}