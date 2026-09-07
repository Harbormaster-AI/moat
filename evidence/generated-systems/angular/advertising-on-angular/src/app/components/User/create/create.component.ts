import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../../../services/User.service';
import { User } from '../../../models/User';
import { SubBaseComponent } from '../../User/sub.base.component';

@Component({
    selector: 'app-create-user',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateUserComponent extends SubBaseComponent implements OnInit {

    title = 'Add User';

    userForm: FormGroup;
    user: User;

    constructor( http: HttpClient,
        private userService: UserService,
        private fb: FormBuilder,
        private router: Router
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

    
    addUser(firstName, lastName, email, Agency, Teams, AdAccounts, Role): void {
        this.userService
        .addUser(firstName, lastName, email, Agency, Teams, AdAccounts, Role)
            .subscribe(() => {
                this.router.navigate(['/indexUser']);
            });
    }

    ngOnInit(): void {
    }
}