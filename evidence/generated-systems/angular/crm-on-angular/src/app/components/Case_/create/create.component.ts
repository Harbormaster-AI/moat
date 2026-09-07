import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Case_Service } from '../../../services/Case_.service';
import { Case_ } from '../../../models/Case_';
import { SubBaseComponent } from '../../Case_/sub.base.component';

@Component({
    selector: 'app-create-case_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCase_Component extends SubBaseComponent implements OnInit {

    title = 'Add Case_';

    case_Form: FormGroup;
    case_: Case_;

    constructor( http: HttpClient,
        private case_Service: Case_Service,
        private fb: FormBuilder,
        private router: Router
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

    
    addCase_(caseNumber, subject, description, slaDue, Organization, Account, Contact, Owner, Team, Activities, CaseComments, Emails, RelatedOpportunities, Status, Priority, Origin, Severity): void {
        this.case_Service
        .addCase_(caseNumber, subject, description, slaDue, Organization, Account, Contact, Owner, Team, Activities, CaseComments, Emails, RelatedOpportunities, Status, Priority, Origin, Severity)
            .subscribe(() => {
                this.router.navigate(['/indexCase_']);
            });
    }

    ngOnInit(): void {
    }
}