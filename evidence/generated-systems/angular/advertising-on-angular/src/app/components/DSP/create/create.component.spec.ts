
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDSPComponent } from './create.component';
import { DSPService } from '../../../services/DSP.service';
import { Router } from '@angular/router';

describe('CreateDSPComponent', () => {
  let component: CreateDSPComponent;
  let fixture: ComponentFixture<CreateDSPComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDSPComponent
      ],
      providers: [
        DSPService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDSPComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});