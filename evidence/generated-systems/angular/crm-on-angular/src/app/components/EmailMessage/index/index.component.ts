
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EmailMessageService } from '../../../services/EmailMessage.service';
import { EmailMessage } from '../../../models/EmailMessage';

@Component({
    selector: 'app-index-emailMessage',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEmailMessageComponent implements OnInit {

    emailMessages: EmailMessage[] = [];

    constructor(
        private router: Router,
        private service: EmailMessageService
) {}

    ngOnInit(): void {
        this.getEmailMessages();
}

    getEmailMessages(): void {
        this.service.getEmailMessages().subscribe((res) => {
        this.emailMessages = res;
    });
}

    deleteEmailMessage(id: any): void {
        this.service.deleteEmailMessage(id)
            .subscribe(() => {
                this.getEmailMessages();
            });
    }
}