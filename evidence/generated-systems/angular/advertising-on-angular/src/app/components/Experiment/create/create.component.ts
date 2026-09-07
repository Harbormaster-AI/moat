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
      hypothesis: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      Campaign: ['', ],
      Variants: ['', ],
      Status: ['', ]
        });
    }

    
    addExperiment(name, hypothesis, startDate, endDate, Campaign, Variants, Status): void {
        this.experimentService
        .addExperiment(name, hypothesis, startDate, endDate, Campaign, Variants, Status)
            .subscribe(() => {
                this.router.navigate(['/indexExperiment']);
            });
    }

    ngOnInit(): void {
    }
}