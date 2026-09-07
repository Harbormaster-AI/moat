
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateClaimReserveComponent } from './create.component';
import { ClaimReserveService } from '../../../services/ClaimReserve.service';
import { Router } from '@angular/router';

describe('CreateClaimReserveComponent', () => {
  let component: CreateClaimReserveComponent;
  let fixture: ComponentFixture<CreateClaimReserveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateClaimReserveComponent
      ],
      providers: [
        ClaimReserveService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateClaimReserveComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});