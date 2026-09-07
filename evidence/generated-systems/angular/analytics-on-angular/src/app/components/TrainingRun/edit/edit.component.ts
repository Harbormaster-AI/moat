import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TrainingRunService } from '../../../services/TrainingRun.service';
import { SubBaseComponent } from '../../TrainingRun/sub.base.component';


@Component({
    selector: 'app-edit-trainingRun',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTrainingRunComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TrainingRun';

    trainingRunForm: FormGroup;
    trainingRun: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TrainingRunService,
        private fb: FormBuilder
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

    
    updateTrainingRun(runLabel, startedAt, completedAt, Experiment, ModelVersion, InputDatasets, Features, RunMetrics, RunParameters, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTrainingRun(runLabel, startedAt, completedAt, Experiment, ModelVersion, InputDatasets, Features, RunMetrics, RunParameters, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTrainingRun']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTrainingRun(params['id']).subscribe(res => {
                this.trainingRun = res;
            });
        });
    }
}