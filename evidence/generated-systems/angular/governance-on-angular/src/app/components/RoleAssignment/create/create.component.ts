import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RoleAssignmentService } from '../../../services/RoleAssignment.service';
import { RoleAssignment } from '../../../models/RoleAssignment';
import { SubBaseComponent } from '../../RoleAssignment/sub.base.component';

@Component({
    selector: 'app-create-roleAssignment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRoleAssignmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add RoleAssignment';

    roleAssignmentForm: FormGroup;
    roleAssignment: RoleAssignment;

    constructor( http: HttpClient,
        private roleAssignmentService: RoleAssignmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.roleAssignmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  effectiveFrom: ['', Validators.required],
      effectiveTo: ['', Validators.required],
      Person: ['', ],
      Role: ['', ],
      GovernanceBody: ['', ],
      Organization: ['', ]
        });
    }

    
    addRoleAssignment(effectiveFrom, effectiveTo, Person, Role, GovernanceBody, Organization): void {
        this.roleAssignmentService
        .addRoleAssignment(effectiveFrom, effectiveTo, Person, Role, GovernanceBody, Organization)
            .subscribe(() => {
                this.router.navigate(['/indexRoleAssignment']);
            });
    }

    ngOnInit(): void {
    }
}