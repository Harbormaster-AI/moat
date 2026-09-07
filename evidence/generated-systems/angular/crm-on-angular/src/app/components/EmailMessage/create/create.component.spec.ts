
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEmailMessageComponent } from './create.component';
import { EmailMessageService } from '../../../services/EmailMessage.service';
import { Router } from '@angular/router';

describe('CreateEmailMessageComponent', () => {
  let component: CreateEmailMessageComponent;
  let fixture: ComponentFixture<CreateEmailMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEmailMessageComponent
      ],
      providers: [
        EmailMessageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEmailMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});