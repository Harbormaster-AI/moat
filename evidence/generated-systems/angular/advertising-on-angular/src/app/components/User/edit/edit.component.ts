import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { UserService } from '../../../services/User.service';
import { SubBaseComponent } from '../../User/sub.base.component';


@Component({
    selector: 'app-edit-user',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditUserComponent extends SubBaseComponent implements OnInit {

    title = 'Edit User';

    userForm: FormGroup;
    user: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: UserService,
        private fb: FormBuilder
) {
        super(http);
        this.userForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      email: ['', Validators.required],
      Agency: ['', ],
      Teams: ['', ],
      AdAccounts: ['', ],
      Role: ['', ]
        });
    }

    
    updateUser(firstName, lastName, email, Agency, Teams, AdAccounts, Role): void {
        this.route.params.subscribe((params) => {

                        this.service.updateUser(firstName, lastName, email, Agency, Teams, AdAccounts, Role, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexUser']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getUser(params['id']).subscribe(res => {
                this.user = res;
            });
        });
    }
}