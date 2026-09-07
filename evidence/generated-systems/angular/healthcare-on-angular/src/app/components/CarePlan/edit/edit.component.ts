import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CarePlanService } from '../../../services/CarePlan.service';
import { SubBaseComponent } from '../../CarePlan/sub.base.component';


@Component({
    selector: 'app-edit-carePlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCarePlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CarePlan';

    carePlanForm: FormGroup;
    carePlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CarePlanService,
        private fb: FormBuilder
) {
        super(http);
        this.carePlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  planNumber: ['', Validators.required],
      goalSummary: ['', Validators.required],
      Patient: ['', ],
      Encounters: ['', ],
      Tasks: ['', ],
      CareTeam: ['', ],
      Status: ['', ]
        });
    }

    
    updateCarePlan(planNumber, goalSummary, Patient, Encounters, Tasks, CareTeam, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCarePlan(planNumber, goalSummary, Patient, Encounters, Tasks, CareTeam, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCarePlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCarePlan(params['id']).subscribe(res => {
                this.carePlan = res;
            });
        });
    }
}