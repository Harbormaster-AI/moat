
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TerminalService } from '../../../services/Terminal.service';
import { Terminal } from '../../../models/Terminal';

@Component({
    selector: 'app-index-terminal',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTerminalComponent implements OnInit {

    terminals: Terminal[] = [];

    constructor(
        private router: Router,
        private service: TerminalService
) {}

    ngOnInit(): void {
        this.getTerminals();
}

    getTerminals(): void {
        this.service.getTerminals().subscribe((res) => {
        this.terminals = res;
    });
}

    deleteTerminal(id: any): void {
        this.service.deleteTerminal(id)
            .subscribe(() => {
                this.getTerminals();
            });
    }
}