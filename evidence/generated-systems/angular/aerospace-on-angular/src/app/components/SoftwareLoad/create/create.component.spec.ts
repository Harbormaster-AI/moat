
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSoftwareLoadComponent } from './create.component';
import { SoftwareLoadService } from '../../../services/SoftwareLoad.service';
import { Router } from '@angular/router';

describe('CreateSoftwareLoadComponent', () => {
  let component: CreateSoftwareLoadComponent;
  let fixture: ComponentFixture<CreateSoftwareLoadComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSoftwareLoadComponent
      ],
      providers: [
        SoftwareLoadService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSoftwareLoadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});