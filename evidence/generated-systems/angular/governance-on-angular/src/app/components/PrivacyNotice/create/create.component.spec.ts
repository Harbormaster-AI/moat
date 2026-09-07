
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePrivacyNoticeComponent } from './create.component';
import { PrivacyNoticeService } from '../../../services/PrivacyNotice.service';
import { Router } from '@angular/router';

describe('CreatePrivacyNoticeComponent', () => {
  let component: CreatePrivacyNoticeComponent;
  let fixture: ComponentFixture<CreatePrivacyNoticeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePrivacyNoticeComponent
      ],
      providers: [
        PrivacyNoticeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePrivacyNoticeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});