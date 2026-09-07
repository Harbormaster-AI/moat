import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RoleService } from '../../../services/Role.service';
import { SubBaseComponent } from '../../Role/sub.base.component';


@Component({
    selector: 'app-edit-role',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRoleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Role';

    roleForm: FormGroup;
    role: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RoleService,
        private fb: FormBuilder
) {
        super(http);
        this.roleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      responsibility: ['', Validators.required],
      Assignments: ['', ]
        });
    }

    
    updateRole(name, responsibility, Assignments): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRole(name, responsibility, Assignments, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRole']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRole(params['id']).subscribe(res => {
                this.role = res;
            });
        });
    }
}