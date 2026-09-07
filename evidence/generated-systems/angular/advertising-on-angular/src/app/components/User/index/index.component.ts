
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../../../services/User.service';
import { User } from '../../../models/User';

@Component({
    selector: 'app-index-user',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexUserComponent implements OnInit {

    users: User[] = [];

    constructor(
        private router: Router,
        private service: UserService
) {}

    ngOnInit(): void {
        this.getUsers();
}

    getUsers(): void {
        this.service.getUsers().subscribe((res) => {
        this.users = res;
    });
}

    deleteUser(id: any): void {
        this.service.deleteUser(id)
            .subscribe(() => {
                this.getUsers();
            });
    }
}