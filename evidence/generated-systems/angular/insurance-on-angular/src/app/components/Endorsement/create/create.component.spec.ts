
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEndorsementComponent } from './create.component';
import { EndorsementService } from '../../../services/Endorsement.service';
import { Router } from '@angular/router';

describe('CreateEndorsementComponent', () => {
  let component: CreateEndorsementComponent;
  let fixture: ComponentFixture<CreateEndorsementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEndorsementComponent
      ],
      providers: [
        EndorsementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEndorsementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});