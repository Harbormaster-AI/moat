
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBillingProfileComponent } from './create.component';
import { BillingProfileService } from '../../../services/BillingProfile.service';
import { Router } from '@angular/router';

describe('CreateBillingProfileComponent', () => {
  let component: CreateBillingProfileComponent;
  let fixture: ComponentFixture<CreateBillingProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBillingProfileComponent
      ],
      providers: [
        BillingProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBillingProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});