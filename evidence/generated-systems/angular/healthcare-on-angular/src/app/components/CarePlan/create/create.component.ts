import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CarePlanService } from '../../../services/CarePlan.service';
import { CarePlan } from '../../../models/CarePlan';
import { SubBaseComponent } from '../../CarePlan/sub.base.component';

@Component({
    selector: 'app-create-carePlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCarePlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add CarePlan';

    carePlanForm: FormGroup;
    carePlan: CarePlan;

    constructor( http: HttpClient,
        private carePlanService: CarePlanService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCarePlan(planNumber, goalSummary, Patient, Encounters, Tasks, CareTeam, Status): void {
        this.carePlanService
        .addCarePlan(planNumber, goalSummary, Patient, Encounters, Tasks, CareTeam, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCarePlan']);
            });
    }

    ngOnInit(): void {
    }
}