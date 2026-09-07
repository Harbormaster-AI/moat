
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSecurityComponent } from './create.component';
import { SecurityService } from '../../../services/Security.service';
import { Router } from '@angular/router';

describe('CreateSecurityComponent', () => {
  let component: CreateSecurityComponent;
  let fixture: ComponentFixture<CreateSecurityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSecurityComponent
      ],
      providers: [
        SecurityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSecurityComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});