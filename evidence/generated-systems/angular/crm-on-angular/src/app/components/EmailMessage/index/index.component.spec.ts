
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEmailMessageComponent } from './index.component';
import { EmailMessageService } from '../../../services/EmailMessage.service';

describe('IndexEmailMessageComponent', () => {
  let component: IndexEmailMessageComponent;
  let fixture: ComponentFixture<IndexEmailMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEmailMessageComponent
      ],
      providers: [
        EmailMessageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEmailMessageComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});