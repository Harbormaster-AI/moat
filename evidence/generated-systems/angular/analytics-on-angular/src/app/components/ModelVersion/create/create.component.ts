import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModelVersionService } from '../../../services/ModelVersion.service';
import { ModelVersion } from '../../../models/ModelVersion';
import { SubBaseComponent } from '../../ModelVersion/sub.base.component';

@Component({
    selector: 'app-create-modelVersion',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateModelVersionComponent extends SubBaseComponent implements OnInit {

    title = 'Add ModelVersion';

    modelVersionForm: FormGroup;
    modelVersion: ModelVersion;

    constructor( http: HttpClient,
        private modelVersionService: ModelVersionService,
        private fb: FormBuilder,
        private router: Router
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

    
    addModelVersion(version, Model_, TrainingRun, EvaluationMetrics, Deployments, FeatureSets, Datasets, Lifecycle, TrainingStatus): void {
        this.modelVersionService
        .addModelVersion(version, Model_, TrainingRun, EvaluationMetrics, Deployments, FeatureSets, Datasets, Lifecycle, TrainingStatus)
            .subscribe(() => {
                this.router.navigate(['/indexModelVersion']);
            });
    }

    ngOnInit(): void {
    }
}