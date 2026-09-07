import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CoverageService } from '../../../services/Coverage.service';
import { Coverage } from '../../../models/Coverage';
import { SubBaseComponent } from '../../Coverage/sub.base.component';

@Component({
    selector: 'app-create-coverage',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCoverageComponent extends SubBaseComponent implements OnInit {

    title = 'Add Coverage';

    coverageForm: FormGroup;
    coverage: Coverage;

    constructor( http: HttpClient,
        private coverageService: CoverageService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.coverageForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  memberId: ['', Validators.required],
      groupNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      endDate: ['', Validators.required],
      Patient: ['', ],
      Plan: ['', ],
      Claims: ['', ],
      Authorizations: ['', ],
      CoverageType: ['', ]
        });
    }

    
    addCoverage(memberId, groupNumber, effectiveDate, endDate, Patient, Plan, Claims, Authorizations, CoverageType): void {
        this.coverageService
        .addCoverage(memberId, groupNumber, effectiveDate, endDate, Patient, Plan, Claims, Authorizations, CoverageType)
            .subscribe(() => {
                this.router.navigate(['/indexCoverage']);
            });
    }

    ngOnInit(): void {
    }
}