
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateExpirationPolicyComponent } from './create.component';
import { ExpirationPolicyService } from '../../../services/ExpirationPolicy.service';
import { Router } from '@angular/router';

describe('CreateExpirationPolicyComponent', () => {
  let component: CreateExpirationPolicyComponent;
  let fixture: ComponentFixture<CreateExpirationPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateExpirationPolicyComponent
      ],
      providers: [
        ExpirationPolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateExpirationPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});