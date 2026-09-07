import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MetricService } from '../../../services/Metric.service';
import { SubBaseComponent } from '../../Metric/sub.base.component';


@Component({
    selector: 'app-edit-metric',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Metric';

    metricForm: FormGroup;
    metric: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MetricService,
        private fb: FormBuilder
) {
        super(http);
        this.metricForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      expression: ['', Validators.required],
      unit: ['', Validators.required],
      SemanticModel: ['', ],
      Datasets: ['', ],
      GlossaryTerms: ['', ],
      Alerts: ['', ],
      Visualizations: ['', ],
      MetricType: ['', ]
        });
    }

    
    updateMetric(name, expression, unit, SemanticModel, Datasets, GlossaryTerms, Alerts, Visualizations, MetricType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMetric(name, expression, unit, SemanticModel, Datasets, GlossaryTerms, Alerts, Visualizations, MetricType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMetric']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMetric(params['id']).subscribe(res => {
                this.metric = res;
            });
        });
    }
}