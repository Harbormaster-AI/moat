import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RoleService } from '../../../services/Role.service';
import { Role } from '../../../models/Role';
import { SubBaseComponent } from '../../Role/sub.base.component';

@Component({
    selector: 'app-create-role',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRoleComponent extends SubBaseComponent implements OnInit {

    title = 'Add Role';

    roleForm: FormGroup;
    role: Role;

    constructor( http: HttpClient,
        private roleService: RoleService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRole(name, responsibility, Assignments): void {
        this.roleService
        .addRole(name, responsibility, Assignments)
            .subscribe(() => {
                this.router.navigate(['/indexRole']);
            });
    }

    ngOnInit(): void {
    }
}