import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ThirdPartyAssessmentService } from '../../../services/ThirdPartyAssessment.service';
import { ThirdPartyAssessment } from '../../../models/ThirdPartyAssessment';
import { SubBaseComponent } from '../../ThirdPartyAssessment/sub.base.component';

@Component({
    selector: 'app-create-thirdPartyAssessment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateThirdPartyAssessmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add ThirdPartyAssessment';

    thirdPartyAssessmentForm: FormGroup;
    thirdPartyAssessment: ThirdPartyAssessment;

    constructor( http: HttpClient,
        private thirdPartyAssessmentService: ThirdPartyAssessmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.thirdPartyAssessmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  assessmentDate: ['', Validators.required],
      assessor: ['', Validators.required],
      ThirdParty: ['', ],
      Issues: ['', ],
      AssessmentType: ['', ],
      Result: ['', ]
        });
    }

    
    addThirdPartyAssessment(assessmentDate, assessor, ThirdParty, Issues, AssessmentType, Result): void {
        this.thirdPartyAssessmentService
        .addThirdPartyAssessment(assessmentDate, assessor, ThirdParty, Issues, AssessmentType, Result)
            .subscribe(() => {
                this.router.navigate(['/indexThirdPartyAssessment']);
            });
    }

    ngOnInit(): void {
    }
}