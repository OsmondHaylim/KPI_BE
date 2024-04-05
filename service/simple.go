package service

import (
	"goreact/model"
)

func (cs *crudService) AddAttendance(input *model.Attendance) error{return cs.attendanceRepo.Store(input)}
func (cs *crudService) AddAnalisa(input *model.Analisa) error{return cs.analisaRepo.Store(input)}
func (cs *crudService) AddFactor(input *model.Factor) error{return cs.factorRepo.Store(input)}
func (cs *crudService) AddFile(input *model.UploadFile) error{return cs.fileRepo.Store(input)}
func (cs *crudService) AddItem(input *model.Item) error{return cs.itemRepo.Store(input)}
func (cs *crudService) AddMasalah(input *model.Masalah) error{return cs.masalahRepo.Store(input)}
func (cs *crudService) AddMinipap(input *model.MiniPAP) error{return cs.minipapRepo.Store(input)}
func (cs *crudService) AddMonthly(input *model.Monthly) error{return cs.monthlyRepo.Store(input)}
func (cs *crudService) AddProject(input *model.Project) error{return cs.projectRepo.Store(input)}
func (cs *crudService) AddResult(input *model.Result) error{return cs.resultRepo.Store(input)}
func (cs *crudService) AddSummary(input *model.Summary) error{return cs.summaryRepo.Store(input)}
func (cs *crudService) AddYearly(input *model.Yearly) error{return cs.yearlyRepo.Store(input)}

func (cs *crudService) UpdateAttendance(input model.Attendance) error{return cs.attendanceRepo.Saves(input)}
func (cs *crudService) UpdateAnalisa(input model.Analisa) error{return cs.analisaRepo.Saves(input)}
func (cs *crudService) UpdateFactor(input model.Factor) error{return cs.factorRepo.Saves(input)}
func (cs *crudService) UpdateFile(input model.UploadFile) error{return cs.fileRepo.Saves(input)}
func (cs *crudService) UpdateItem(input model.Item) error{return cs.itemRepo.Saves(input)}
func (cs *crudService) UpdateMasalah(input model.Masalah) error{return cs.masalahRepo.Saves(input)}
func (cs *crudService) UpdateMinipap(input model.MiniPAP) error{return cs.minipapRepo.Saves(input)}
func (cs *crudService) UpdateMonthly(input model.Monthly) error{return cs.monthlyRepo.Saves(input)}
func (cs *crudService) UpdateProject(input model.Project) error{return cs.projectRepo.Saves(input)}
func (cs *crudService) UpdateResult(input model.Result) error{return cs.resultRepo.Saves(input)}
func (cs *crudService) UpdateSummary(input model.Summary) error{return cs.summaryRepo.Saves(input)}
func (cs *crudService) UpdateYearly(input model.Yearly) error{return cs.yearlyRepo.Saves(input)}

func (cs *crudService) DeleteAttendance(input int) error{return cs.attendanceRepo.Delete(input)}
func (cs *crudService) DeleteAnalisa(input int) error{return cs.analisaRepo.Delete(input)}
func (cs *crudService) DeleteFactor(input int) error{return cs.factorRepo.Delete(input)}
func (cs *crudService) DeleteFile(input int) error{return cs.fileRepo.Delete(input)}
func (cs *crudService) DeleteItem(input int) error{return cs.itemRepo.Delete(input)}
func (cs *crudService) DeleteMasalah(input int) error{return cs.masalahRepo.Delete(input)}
func (cs *crudService) DeleteMinipap(input int) error{return cs.minipapRepo.Delete(input)}
func (cs *crudService) DeleteMonthly(input int) error{return cs.monthlyRepo.Delete(input)}
func (cs *crudService) DeleteProject(input int) error{return cs.projectRepo.Delete(input)}
func (cs *crudService) DeleteResult(input int) error{return cs.resultRepo.Delete(input)}
func (cs *crudService) DeleteSummary(input int) error{return cs.summaryRepo.Delete(input)}
func (cs *crudService) DeleteYearly(input int) error{return cs.yearlyRepo.Delete(input)}

func (cs *crudService) DeleteAttendance(input int) error{return cs.attendanceRepo.Delete(input)}
func (cs *crudService) DeleteAnalisa(input int) error{return cs.analisaRepo.Delete(input)}
func (cs *crudService) DeleteFactor(input int) error{return cs.factorRepo.Delete(input)}
func (cs *crudService) DeleteFile(input int) error{return cs.fileRepo.Delete(input)}
func (cs *crudService) DeleteItem(input int) error{return cs.itemRepo.Delete(input)}
func (cs *crudService) DeleteMasalah(input int) error{return cs.masalahRepo.Delete(input)}
func (cs *crudService) DeleteMinipap(input int) error{return cs.minipapRepo.Delete(input)}
func (cs *crudService) DeleteMonthly(input int) error{return cs.monthlyRepo.Delete(input)}
func (cs *crudService) DeleteProject(input int) error{return cs.projectRepo.Delete(input)}
func (cs *crudService) DeleteResult(input int) error{return cs.resultRepo.Delete(input)}
func (cs *crudService) DeleteSummary(input int) error{return cs.summaryRepo.Delete(input)}
func (cs *crudService) DeleteYearly(input int) error{return cs.yearlyRepo.Delete(input)}