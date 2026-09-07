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
                  username: ['', Validators.required],
      fullName: ['', Validators.required],
      email: ['', Validators.required],
      locale: ['', Validators.required],
      Organization: ['', ],
      Teams: ['', ],
      Activities: ['', ],
      OwnedAccounts: ['', ],
      OwnedLeads: ['', ],
      OwnedOpportunities: ['', ],
      OwnedCases: ['', ],
      Quotes: ['', ],
      Orders: ['', ],
      Contracts: ['', ],
      EmailMessages: ['', ],
      Role: ['', ],
      Status: ['', ]
        });
    }

    
    addUser(username, fullName, email, locale, Organization, Teams, Activities, OwnedAccounts, OwnedLeads, OwnedOpportunities, OwnedCases, Quotes, Orders, Contracts, EmailMessages, Role, Status): void {
        this.userService
        .addUser(username, fullName, email, locale, Organization, Teams, Activities, OwnedAccounts, OwnedLeads, OwnedOpportunities, OwnedCases, Quotes, Orders, Contracts, EmailMessages, Role, Status)
            .subscribe(() => {
                this.router.navigate(['/indexUser']);
            });
    }

    ngOnInit(): void {
    }
}