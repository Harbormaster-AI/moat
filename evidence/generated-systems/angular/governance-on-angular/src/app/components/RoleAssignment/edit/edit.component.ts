import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RoleAssignmentService } from '../../../services/RoleAssignment.service';
import { SubBaseComponent } from '../../RoleAssignment/sub.base.component';


@Component({
    selector: 'app-edit-roleAssignment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRoleAssignmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RoleAssignment';

    roleAssignmentForm: FormGroup;
    roleAssignment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RoleAssignmentService,
        private fb: FormBuilder
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

    
    updateRoleAssignment(effectiveFrom, effectiveTo, Person, Role, GovernanceBody, Organization): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRoleAssignment(effectiveFrom, effectiveTo, Person, Role, GovernanceBody, Organization, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRoleAssignment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRoleAssignment(params['id']).subscribe(res => {
                this.roleAssignment = res;
            });
        });
    }
}