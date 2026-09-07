
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PersonService } from '../../../services/Person.service';
import { Person } from '../../../models/Person';

@Component({
    selector: 'app-index-person',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPersonComponent implements OnInit {

    persons: Person[] = [];

    constructor(
        private router: Router,
        private service: PersonService
) {}

    ngOnInit(): void {
        this.getPersons();
}

    getPersons(): void {
        this.service.getPersons().subscribe((res) => {
        this.persons = res;
    });
}

    deletePerson(id: any): void {
        this.service.deletePerson(id)
            .subscribe(() => {
                this.getPersons();
            });
    }
}