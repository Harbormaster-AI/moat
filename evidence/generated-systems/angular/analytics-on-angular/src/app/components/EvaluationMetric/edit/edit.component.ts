import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EvaluationMetricService } from '../../../services/EvaluationMetric.service';
import { SubBaseComponent } from '../../EvaluationMetric/sub.base.component';


@Component({
    selector: 'app-edit-evaluationMetric',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEvaluationMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Edit EvaluationMetric';

    evaluationMetricForm: FormGroup;
    evaluationMetric: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EvaluationMetricService,
        private fb: FormBuilder
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

    
    updateEvaluationMetric(name, value, ModelVersion, Metric, Dataset): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEvaluationMetric(name, value, ModelVersion, Metric, Dataset, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEvaluationMetric']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEvaluationMetric(params['id']).subscribe(res => {
                this.evaluationMetric = res;
            });
        });
    }
}