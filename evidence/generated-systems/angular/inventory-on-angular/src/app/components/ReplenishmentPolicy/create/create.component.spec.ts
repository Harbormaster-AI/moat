
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateReplenishmentPolicyComponent } from './create.component';
import { ReplenishmentPolicyService } from '../../../services/ReplenishmentPolicy.service';
import { Router } from '@angular/router';

describe('CreateReplenishmentPolicyComponent', () => {
  let component: CreateReplenishmentPolicyComponent;
  let fixture: ComponentFixture<CreateReplenishmentPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateReplenishmentPolicyComponent
      ],
      providers: [
        ReplenishmentPolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateReplenishmentPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});