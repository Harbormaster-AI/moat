
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexControlTest_Component } from './index.component';
import { ControlTest_Service } from '../../../services/ControlTest_.service';

describe('IndexControlTest_Component', () => {
  let component: IndexControlTest_Component;
  let fixture: ComponentFixture<IndexControlTest_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexControlTest_Component
      ],
      providers: [
        ControlTest_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexControlTest_Component);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});