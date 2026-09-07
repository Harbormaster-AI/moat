import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { Case_Service } from '../../../services/Case_.service';
import { SubBaseComponent } from '../../Case_/sub.base.component';


@Component({
    selector: 'app-edit-case_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCase_Component extends SubBaseComponent implements OnInit {

    title = 'Edit Case_';

    case_Form: FormGroup;
    case_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: Case_Service,
        private fb: FormBuilder
) {
        super(http);
        this.case_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  caseNumber: ['', Validators.required],
      subject: ['', Validators.required],
      description: ['', Validators.required],
      slaDue: ['', Validators.required],
      Organization: ['', ],
      Account: ['', ],
      Contact: ['', ],
      Owner: ['', ],
      Team: ['', ],
      Activities: ['', ],
      CaseComments: ['', ],
      Emails: ['', ],
      RelatedOpportunities: ['', ],
      Status: ['', ],
      Priority: ['', ],
      Origin: ['', ],
      Severity: ['', ]
        });
    }

    
    updateCase_(caseNumber, subject, description, slaDue, Organization, Account, Contact, Owner, Team, Activities, CaseComments, Emails, RelatedOpportunities, Status, Priority, Origin, Severity): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCase_(caseNumber, subject, description, slaDue, Organization, Account, Contact, Owner, Team, Activities, CaseComments, Emails, RelatedOpportunities, Status, Priority, Origin, Severity, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCase_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCase_(params['id']).subscribe(res => {
                this.case_ = res;
            });
        });
    }
}