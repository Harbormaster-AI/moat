import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ExperimentService } from '../../../services/Experiment.service';
import { Experiment } from '../../../models/Experiment';
import { SubBaseComponent } from '../../Experiment/sub.base.component';

@Component({
    selector: 'app-create-experiment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateExperimentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Experiment';

    experimentForm: FormGroup;
    experiment: Experiment;

    constructor( http: HttpClient,
        private experimentService: ExperimentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.experimentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      objective: ['', Validators.required],
      Workspace: ['', ],
      TrainingRuns: ['', ],
      Models: ['', ],
      Notebooks: ['', ],
      Status: ['', ]
        });
    }

    
    addExperiment(name, objective, Workspace, TrainingRuns, Models, Notebooks, Status): void {
        this.experimentService
        .addExperiment(name, objective, Workspace, TrainingRuns, Models, Notebooks, Status)
            .subscribe(() => {
                this.router.navigate(['/indexExperiment']);
            });
    }

    ngOnInit(): void {
    }
}