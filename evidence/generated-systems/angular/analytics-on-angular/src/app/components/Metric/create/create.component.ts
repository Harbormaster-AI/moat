import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MetricService } from '../../../services/Metric.service';
import { Metric } from '../../../models/Metric';
import { SubBaseComponent } from '../../Metric/sub.base.component';

@Component({
    selector: 'app-create-metric',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Add Metric';

    metricForm: FormGroup;
    metric: Metric;

    constructor( http: HttpClient,
        private metricService: MetricService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMetric(name, expression, unit, SemanticModel, Datasets, GlossaryTerms, Alerts, Visualizations, MetricType): void {
        this.metricService
        .addMetric(name, expression, unit, SemanticModel, Datasets, GlossaryTerms, Alerts, Visualizations, MetricType)
            .subscribe(() => {
                this.router.navigate(['/indexMetric']);
            });
    }

    ngOnInit(): void {
    }
}