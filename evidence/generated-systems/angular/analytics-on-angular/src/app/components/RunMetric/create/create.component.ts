import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RunMetricService } from '../../../services/RunMetric.service';
import { RunMetric } from '../../../models/RunMetric';
import { SubBaseComponent } from '../../RunMetric/sub.base.component';

@Component({
    selector: 'app-create-runMetric',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRunMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Add RunMetric';

    runMetricForm: FormGroup;
    runMetric: RunMetric;

    constructor( http: HttpClient,
        private runMetricService: RunMetricService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.runMetricForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      value: ['', Validators.required],
      TrainingRun: ['', ],
      Metric: ['', ],
      Dataset: ['', ]
        });
    }

    
    addRunMetric(name, value, TrainingRun, Metric, Dataset): void {
        this.runMetricService
        .addRunMetric(name, value, TrainingRun, Metric, Dataset)
            .subscribe(() => {
                this.router.navigate(['/indexRunMetric']);
            });
    }

    ngOnInit(): void {
    }
}