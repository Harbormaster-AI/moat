import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { GovernanceBodyService } from '../../../services/GovernanceBody.service';
import { GovernanceBody } from '../../../models/GovernanceBody';
import { SubBaseComponent } from '../../GovernanceBody/sub.base.component';

@Component({
    selector: 'app-create-governanceBody',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateGovernanceBodyComponent extends SubBaseComponent implements OnInit {

    title = 'Add GovernanceBody';

    governanceBodyForm: FormGroup;
    governanceBody: GovernanceBody;

    constructor( http: HttpClient,
        private governanceBodyService: GovernanceBodyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.governanceBodyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      charterUrl: ['', Validators.required],
      chair: ['', Validators.required],
      Organization: ['', ],
      RoleAssignments: ['', ],
      Policies: ['', ],
      BodyType: ['', ]
        });
    }

    
    addGovernanceBody(name, charterUrl, chair, Organization, RoleAssignments, Policies, BodyType): void {
        this.governanceBodyService
        .addGovernanceBody(name, charterUrl, chair, Organization, RoleAssignments, Policies, BodyType)
            .subscribe(() => {
                this.router.navigate(['/indexGovernanceBody']);
            });
    }

    ngOnInit(): void {
    }
}