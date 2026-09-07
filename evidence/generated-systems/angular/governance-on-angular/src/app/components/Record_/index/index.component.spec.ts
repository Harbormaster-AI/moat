
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRecord_Component } from './index.component';
import { Record_Service } from '../../../services/Record_.service';

describe('IndexRecord_Component', () => {
  let component: IndexRecord_Component;
  let fixture: ComponentFixture<IndexRecord_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRecord_Component
      ],
      providers: [
        Record_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRecord_Component);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});