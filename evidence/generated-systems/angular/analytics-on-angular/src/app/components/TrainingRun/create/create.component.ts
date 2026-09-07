import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TrainingRunService } from '../../../services/TrainingRun.service';
import { TrainingRun } from '../../../models/TrainingRun';
import { SubBaseComponent } from '../../TrainingRun/sub.base.component';

@Component({
    selector: 'app-create-trainingRun',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTrainingRunComponent extends SubBaseComponent implements OnInit {

    title = 'Add TrainingRun';

    trainingRunForm: FormGroup;
    trainingRun: TrainingRun;

    constructor( http: HttpClient,
        private trainingRunService: TrainingRunService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.trainingRunForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  runLabel: ['', Validators.required],
      startedAt: ['', Validators.required],
      completedAt: ['', Validators.required],
      Experiment: ['', ],
      ModelVersion: ['', ],
      InputDatasets: ['', ],
      Features: ['', ],
      RunMetrics: ['', ],
      RunParameters: ['', ],
      Status: ['', ]
        });
    }

    
    addTrainingRun(runLabel, startedAt, completedAt, Experiment, ModelVersion, InputDatasets, Features, RunMetrics, RunParameters, Status): void {
        this.trainingRunService
        .addTrainingRun(runLabel, startedAt, completedAt, Experiment, ModelVersion, InputDatasets, Features, RunMetrics, RunParameters, Status)
            .subscribe(() => {
                this.router.navigate(['/indexTrainingRun']);
            });
    }

    ngOnInit(): void {
    }
}