import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { GovernanceBodyService } from '../../../services/GovernanceBody.service';
import { SubBaseComponent } from '../../GovernanceBody/sub.base.component';


@Component({
    selector: 'app-edit-governanceBody',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditGovernanceBodyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit GovernanceBody';

    governanceBodyForm: FormGroup;
    governanceBody: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: GovernanceBodyService,
        private fb: FormBuilder
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

    
    updateGovernanceBody(name, charterUrl, chair, Organization, RoleAssignments, Policies, BodyType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateGovernanceBody(name, charterUrl, chair, Organization, RoleAssignments, Policies, BodyType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexGovernanceBody']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getGovernanceBody(params['id']).subscribe(res => {
                this.governanceBody = res;
            });
        });
    }
}