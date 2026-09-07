import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ThirdPartyAssessmentService } from '../../../services/ThirdPartyAssessment.service';
import { SubBaseComponent } from '../../ThirdPartyAssessment/sub.base.component';


@Component({
    selector: 'app-edit-thirdPartyAssessment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditThirdPartyAssessmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ThirdPartyAssessment';

    thirdPartyAssessmentForm: FormGroup;
    thirdPartyAssessment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ThirdPartyAssessmentService,
        private fb: FormBuilder
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

    
    updateThirdPartyAssessment(assessmentDate, assessor, ThirdParty, Issues, AssessmentType, Result): void {
        this.route.params.subscribe((params) => {

                        this.service.updateThirdPartyAssessment(assessmentDate, assessor, ThirdParty, Issues, AssessmentType, Result, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexThirdPartyAssessment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getThirdPartyAssessment(params['id']).subscribe(res => {
                this.thirdPartyAssessment = res;
            });
        });
    }
}