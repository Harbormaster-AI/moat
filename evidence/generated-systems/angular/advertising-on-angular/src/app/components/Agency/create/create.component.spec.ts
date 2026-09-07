
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAgencyComponent } from './create.component';
import { AgencyService } from '../../../services/Agency.service';
import { Router } from '@angular/router';

describe('CreateAgencyComponent', () => {
  let component: CreateAgencyComponent;
  let fixture: ComponentFixture<CreateAgencyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAgencyComponent
      ],
      providers: [
        AgencyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAgencyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});