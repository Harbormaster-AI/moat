import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InspectionPlanService } from '../../../services/InspectionPlan.service';
import { InspectionPlan } from '../../../models/InspectionPlan';
import { SubBaseComponent } from '../../InspectionPlan/sub.base.component';

@Component({
    selector: 'app-create-inspectionPlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInspectionPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add InspectionPlan';

    inspectionPlanForm: FormGroup;
    inspectionPlan: InspectionPlan;

    constructor( http: HttpClient,
        private inspectionPlanService: InspectionPlanService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inspectionPlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  planNumber: ['', Validators.required],
      revision: ['', Validators.required],
      Item: ['', ],
      Characteristics: ['', ],
      SamplingPlan: ['', ],
      Status: ['', ]
        });
    }

    
    addInspectionPlan(planNumber, revision, Item, Characteristics, SamplingPlan, Status): void {
        this.inspectionPlanService
        .addInspectionPlan(planNumber, revision, Item, Characteristics, SamplingPlan, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInspectionPlan']);
            });
    }

    ngOnInit(): void {
    }
}