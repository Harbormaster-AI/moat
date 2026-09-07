
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AirworthinessDirectiveService } from '../../../services/AirworthinessDirective.service';
import { AirworthinessDirective } from '../../../models/AirworthinessDirective';

@Component({
    selector: 'app-index-airworthinessDirective',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAirworthinessDirectiveComponent implements OnInit {

    airworthinessDirectives: AirworthinessDirective[] = [];

    constructor(
        private router: Router,
        private service: AirworthinessDirectiveService
) {}

    ngOnInit(): void {
        this.getAirworthinessDirectives();
}

    getAirworthinessDirectives(): void {
        this.service.getAirworthinessDirectives().subscribe((res) => {
        this.airworthinessDirectives = res;
    });
}

    deleteAirworthinessDirective(id: any): void {
        this.service.deleteAirworthinessDirective(id)
            .subscribe(() => {
                this.getAirworthinessDirectives();
            });
    }
}