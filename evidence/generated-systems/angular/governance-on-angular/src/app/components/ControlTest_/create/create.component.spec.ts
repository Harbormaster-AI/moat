
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateControlTest_Component } from './create.component';
import { ControlTest_Service } from '../../../services/ControlTest_.service';
import { Router } from '@angular/router';

describe('CreateControlTest_Component', () => {
  let component: CreateControlTest_Component;
  let fixture: ComponentFixture<CreateControlTest_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateControlTest_Component
      ],
      providers: [
        ControlTest_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateControlTest_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});