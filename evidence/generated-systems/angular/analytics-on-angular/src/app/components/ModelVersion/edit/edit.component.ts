import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ModelVersionService } from '../../../services/ModelVersion.service';
import { SubBaseComponent } from '../../ModelVersion/sub.base.component';


@Component({
    selector: 'app-edit-modelVersion',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditModelVersionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ModelVersion';

    modelVersionForm: FormGroup;
    modelVersion: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ModelVersionService,
        private fb: FormBuilder
) {
        super(http);
        this.modelVersionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  version: ['', Validators.required],
      Model_: ['', ],
      TrainingRun: ['', ],
      EvaluationMetrics: ['', ],
      Deployments: ['', ],
      FeatureSets: ['', ],
      Datasets: ['', ],
      Lifecycle: ['', ],
      TrainingStatus: ['', ]
        });
    }

    
    updateModelVersion(version, Model_, TrainingRun, EvaluationMetrics, Deployments, FeatureSets, Datasets, Lifecycle, TrainingStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateModelVersion(version, Model_, TrainingRun, EvaluationMetrics, Deployments, FeatureSets, Datasets, Lifecycle, TrainingStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexModelVersion']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getModelVersion(params['id']).subscribe(res => {
                this.modelVersion = res;
            });
        });
    }
}