
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProcedureComponent } from './index.component';
import { ProcedureService } from '../../../services/Procedure.service';

describe('IndexProcedureComponent', () => {
  let component: IndexProcedureComponent;
  let fixture: ComponentFixture<IndexProcedureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProcedureComponent
      ],
      providers: [
        ProcedureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProcedureComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});