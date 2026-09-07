import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RunMetricService } from '../../../services/RunMetric.service';
import { SubBaseComponent } from '../../RunMetric/sub.base.component';


@Component({
    selector: 'app-edit-runMetric',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRunMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RunMetric';

    runMetricForm: FormGroup;
    runMetric: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RunMetricService,
        private fb: FormBuilder
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

    
    updateRunMetric(name, value, TrainingRun, Metric, Dataset): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRunMetric(name, value, TrainingRun, Metric, Dataset, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRunMetric']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRunMetric(params['id']).subscribe(res => {
                this.runMetric = res;
            });
        });
    }
}