import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BenefitEnrollmentService } from '../../../services/BenefitEnrollment.service';
import { BenefitEnrollment } from '../../../models/BenefitEnrollment';
import { SubBaseComponent } from '../../BenefitEnrollment/sub.base.component';

@Component({
    selector: 'app-create-benefitEnrollment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBenefitEnrollmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add BenefitEnrollment';

    benefitEnrollmentForm: FormGroup;
    benefitEnrollment: BenefitEnrollment;

    constructor( http: HttpClient,
        private benefitEnrollmentService: BenefitEnrollmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.benefitEnrollmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  enrollmentId: ['', Validators.required],
      effectiveFrom: ['', Validators.required],
      effectiveTo: ['', Validators.required],
      BenefitPlan: ['', ],
      Employee: ['', ],
      Dependents: ['', ],
      Status: ['', ],
      CoverageLevel: ['', ]
        });
    }

    
    addBenefitEnrollment(enrollmentId, effectiveFrom, effectiveTo, BenefitPlan, Employee, Dependents, Status, CoverageLevel): void {
        this.benefitEnrollmentService
        .addBenefitEnrollment(enrollmentId, effectiveFrom, effectiveTo, BenefitPlan, Employee, Dependents, Status, CoverageLevel)
            .subscribe(() => {
                this.router.navigate(['/indexBenefitEnrollment']);
            });
    }

    ngOnInit(): void {
    }
}