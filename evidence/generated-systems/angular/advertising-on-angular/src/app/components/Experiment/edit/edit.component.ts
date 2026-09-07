import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ExperimentService } from '../../../services/Experiment.service';
import { SubBaseComponent } from '../../Experiment/sub.base.component';


@Component({
    selector: 'app-edit-experiment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditExperimentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Experiment';

    experimentForm: FormGroup;
    experiment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ExperimentService,
        private fb: FormBuilder
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

    
    updateExperiment(name, hypothesis, startDate, endDate, Campaign, Variants, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateExperiment(name, hypothesis, startDate, endDate, Campaign, Variants, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexExperiment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getExperiment(params['id']).subscribe(res => {
                this.experiment = res;
            });
        });
    }
}